def fat(n):

	if n < 0:
		print "Error!"
		return -1

	if n <= 1:
		return 1

	return n*fat(n-1)

pares = [0,2,4,6,8]
impares = [1,3,5,7,9]
lista = pares + impares

print("Minha lista:")

for cont, x in enumerate(lista):
	string = str(cont) + ") " + str(x)
	print(string)
print("Length = %d" % len(string))

print "\n-------------"
n = input("Type a number: ")
print ("%d! = %d" % (n, fat(n)) )