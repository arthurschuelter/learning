import argparse

def rewind(file):
	file.seek(0)

def main(args):
	file = open(args.file, 'r', encoding = 'utf-8-sig')
	print('Your file is: >' + args.file + '<')
	
	rewind(file)
	wordcount = {}
	text = file.read().split()

	for word in text:
		if word in wordcount:
			wordcount[word]+= 1
		else:
			wordcount[word] = 1

	for word, num in sorted(wordcount.items()):
		print( word + '\t' + str(num) )


if __name__ == "__main__":
	parser = argparse.ArgumentParser(description="Wordcount")
	parser.add_argument('--file', '-f',
						type=str,
						help="Select the file to be analyzed",
						default="teste.txt")

	args = parser.parse_args()

	main(args)