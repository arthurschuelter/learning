function fact (n)
    if n <= 1 then
        return 1
    end
    return n * fact(n-1)
end

for i = 0, 5 do
    print(string.format("%i) Hello World!", i))
end

print("Type a number: ")
n = io.read("*number")
print(string.format("%d! = %d", n, fact(n)))
