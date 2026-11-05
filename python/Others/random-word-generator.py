import random

def main():
	file = open("teste.txt", "w", encoding="utf-8-sig")

	for i in range(0, 10000):
		a = random.randint(0, 5)
		
		if a == 0: file.write("Dog ")
		if a == 1: file.write("Cat ")
		if a == 2: file.write("Bird ")
		if a == 3: file.write("Train ")
		if a == 4: file.write("Bus ")
		if a == 5: file.write("Boat ")

if __name__== '__main__':
	main()