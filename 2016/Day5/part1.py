import hashlib


def main():
    input = 'wtnhxymk'
    index = 0
    password = ''
    while (len(password) < 8):
        code = input + str(index)
        hash = hashlib.md5(code.encode()).hexdigest()
        if (hash[:5] == '00000'):
            password = password + hash[5]
            print('Found ', len(password), ' chars for password')
        index = index + 1
    print('Password: ', password)


main()
