import hashlib
from sys import stdout


def cinematic_decryption(password, hash):
    """ Prints a cinematic 'decryption' of the password """
    plist = list(password)
    for i in range(0, len(plist)):
        if (plist[i] == ' '):
            plist[i] = hash[i]
    cinematic = ''.join(plist)
    stdout.write("\r%s" % cinematic)
    stdout.flush()


def main():
    input = 'wtnhxymk'
    index = 0
    password = ' ' * 8
    while (' ' in password):
        code = input + str(index)
        hash = hashlib.md5(code.encode()).hexdigest()
        if (hash[:5] == '00000'):
            position = hash[5]
            if (position.isdigit()):
                position = int(position)
                if (position in range(0, 8) and (password[position] == ' ')):
                    plist = list(password)
                    plist[position] = hash[6]
                    password = ''.join(plist)
        cinematic_decryption(password, hash)
        index = index + 1
    stdout.write("\n") # move cursor to the next line after cinematic has finished
    print('Password: ', password)


main()
