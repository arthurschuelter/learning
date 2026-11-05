from collections import Counter, deque
from email.mime import text

class BPETokenizer:

    def __init__(self):
        self.vocab = {}
        self.inverse_vocab = {}
        self.bpe_merges = {}
        self.bpe_ranks = {}


    def train(self, text, vocab_size=1000, allowed_special={"<|endoftext|>"}):
        processed_text = []
        for i, char in enumerate(text):
            if char == " " and i != 0:
                processed_text.append("Ġ")
            if char != " ":
                processed_text.append(char)
        processed_text = "".join(processed_text)

        unique_chars = [chr(i) for i in range(256)]
        unique_chars.extend(
            char for char in sorted(set(processed_text))
            if char not in unique_chars
        )
        if "Ġ" not in unique_chars:
            unique_chars.append("Ġ")

        self.vocab = {i: char for i, char in enumerate(unique_chars)}
        self.inverse_vocab = {char: i for i, char in self.vocab.items()}

        if allowed_special:
            for token in allowed_special:
                if token not in self.inverse_vocab:
                    new_id = len(self.vocab)
                    self.vocab[new_id] = token
                    self.inverse_vocab[token] = new_id
        token_ids = [self.inverse_vocab[char] for char in processed_text]

        for new_id in range(len(self.vocab), vocab_size):
            pair_id = self.find_freq_pair(token_ids, mode="most")
            if pair_id is None:
                break
            token_ids = self.replace_pair(token_ids, pair_id, new_id)
            self.bpe_merges[pair_id] = new_id

        # Build the vocabulary with merged tokens
        for (p0, p1), new_id in self.bpe_merges.items():
            merged_token = self.vocab[p0] + self.vocab[p1]
            self.vocab[new_id] = merged_token
            self.inverse_vocab[merged_token] = new_id


    def encode(self, text, allowed_special={"<|endoftext|>"}):
        import re

        token_ids = []

        tokens = []
        lines = text.split("\n")
        for i, line in enumerate(lines):
            if i > 0:
                tokens.append("\n")
            words = line.split()
            for j, word in enumerate(words):
                if j == 0 and i > 0:
                    tokens.append("Ġ" + word)
                elif j == 0:
                    tokens.append(word)
                else:
                    tokens.append("Ġ" + word)
        for token in tokens:
            if token in self.inverse_vocab:
                token_ids.append(self.inverse_vocab[token])
            else:
                token_ids.extend(self.tokenize_with_bpe(token))
        
        return token_ids

    def tokenize_with_bpe(self, token):
        """
        Tokenize a single token using BPE merges.

        Args:
            token (str): The token to tokenize.

        Returns:
            List[int]: The list of token IDs after applying BPE.
        """
        # Tokenize the token into individual characters (as initial token IDs)
        token_ids = [self.inverse_vocab.get(char, None) for char in token]
        if None in token_ids:
            missing_chars = [char for char, tid in zip(token, token_ids) if tid is None]
            raise ValueError(f"Characters not found in vocab: {missing_chars}")

        # If we haven't loaded OpenAI's GPT-2 merges, use my approach
        if not self.bpe_ranks:
            can_merge = True
            while can_merge and len(token_ids) > 1:
                can_merge = False
                new_tokens = []
                i = 0
                while i < len(token_ids) - 1:
                    pair = (token_ids[i], token_ids[i + 1])
                    if pair in self.bpe_merges:
                        merged_token_id = self.bpe_merges[pair]
                        new_tokens.append(merged_token_id)
                        # Uncomment for educational purposes:
                        # print(f"Merged pair {pair} -> {merged_token_id} ('{self.vocab[merged_token_id]}')")
                        i += 2  # Skip the next token as it's merged
                        can_merge = True
                    else:
                        new_tokens.append(token_ids[i])
                        i += 1
                if i < len(token_ids):
                    new_tokens.append(token_ids[i])
                token_ids = new_tokens
            return token_ids

        # Otherwise, do GPT-2-style merging with the ranks:
        # 1) Convert token_ids back to string "symbols" for each ID
        symbols = [self.vocab[id_num] for id_num in token_ids]

        # Repeatedly merge all occurrences of the lowest-rank pair
        while True:
            # Collect all adjacent pairs
            pairs = set(zip(symbols, symbols[1:]))
            if not pairs:
                break

            # Find the pair with the best (lowest) rank
            min_rank = float("inf")
            bigram = None
            for p in pairs:
                r = self.bpe_ranks.get(p, float("inf"))
                if r < min_rank:
                    min_rank = r
                    bigram = p

            # If no valid ranked pair is present, we're done
            if bigram is None or bigram not in self.bpe_ranks:
                break

            # Merge all occurrences of that pair
            first, second = bigram
            new_symbols = []
            i = 0
            while i < len(symbols):
                # If we see (first, second) at position i, merge them
                if i < len(symbols) - 1 and symbols[i] == first and symbols[i+1] == second:
                    new_symbols.append(first + second)  # merged symbol
                    i += 2
                else:
                    new_symbols.append(symbols[i])
                    i += 1
            symbols = new_symbols

            if len(symbols) == 1:
                break

        # Finally, convert merged symbols back to IDs
        merged_ids = [self.inverse_vocab[sym] for sym in symbols]
        return merged_ids

    def decode(self, token_ids):
        decoded_string = ""
        for i, token_id in enumerate(token_ids):
            if token_id not in self.vocab:
                raise ValueError(f"Token ID {token_id} not found in vocab.")
            token = self.vocab[token_id]
            if token == "\n":
                if decoded_string and not decoded_string.endswith(" "):
                    decoded_string += " "  # Add space if not present before a newline
                decoded_string += token
            elif token.startswith("Ġ"):
                decoded_string += " " + token[1:]
            else:
                decoded_string += token
        return decoded_string


    def find_freq_pair(self, token_ids, mode="most"):
        pairs = Counter(zip(token_ids, token_ids[1:]))

        if not pairs:
            return None
        
        if mode == "most":
            return max(pairs.items(), key=lambda x: x[1])[0]
        elif mode == "least":
            return min(pairs.items(), key=lambda x: x[1])[0]
        else:
            raise ValueError("Mode must be 'most' or 'least'")

    def replace_pair(self, token_ids, pair_id, new_id):
        dq = deque(token_ids)
        replaced = []
        
        while dq:
            curent = dq.popleft()
            if dq and (curent, dq[0]) == pair_id:
                replaced.append(new_id)
                dq.popleft()  # Remove the next token as well
            else:
                replaced.append(curent)
        
        return replaced