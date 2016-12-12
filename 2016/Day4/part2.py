def rotate_letter(c, nr):
    code = ord(c) - 97
    for i in range(0, nr):
        if (code == 25):
            code = 0
        else:
            code = code + 1
    return chr(code + 97)


def decrypt(string, id):
    word = ''
    string = ''.join(string)
    for c in string:
        c = rotate_letter(c, id)
        word = word + c
    return word


def main():
    id = 0
    with open('input.txt') as input:
        for line in input.readlines():
            chars = line.strip().split('-')
            code = chars[-1].split('[')[0]
            word = decrypt(chars[:-1], int(code))
            if ('northpole' in word):
                id = code
                break
    print('Sector ID: ', id)


main()
