from collections import Counter, deque
from bpe_tokenizer import BPETokenizer

def main():
    text = "hello, world! This is a tokenizer hell."
    # tokens = simple_tokenize(text)
    # print(tokens)
    bpe_tokenizer = BPETokenizer()
    bpe_tokenizer.train(text)
    encoded = bpe_tokenizer.encode(text)
    print(f"Encoded token IDs: {encoded}")
    print([bpe_tokenizer.vocab[token_id] for token_id in encoded])

    decoded = bpe_tokenizer.decode(encoded)
    print(f"Decoded text: {decoded}")

def simple_tokenize(text):
    return text.split()

if __name__ == "__main__":    
    main()