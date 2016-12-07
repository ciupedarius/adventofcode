def main():
    with open('input.txt') as file:
        pointer = file.tell()  # save initial file pointer
        size = len(file.readline().strip())
        file.seek(pointer)  # go back to the pointer saved
        col = ["" for x in range(size)]
        for line in file.readlines():
            for i, c in enumerate(line.strip()):
                col[i] = col[i] + c
        code = ""
        for i in range(0, len(col)):
            code = code + min([(j, col[i].count(j)) for j in set(col[i])], key=lambda x: x[1])[0]
        print('Code: ', code)


main()
