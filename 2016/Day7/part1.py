import re


def is_palindrome(string):
    pali = False
    for i in range(0, len(string) - 3):
        word = string[i:i + 4]
        if (word == word[::-1]):
            if word[0] == word[1]:
                pali = False
            else:
                pali = True
                break
    return pali


def main():
    pattern = r'(?<=\[).+?(?=\])'
    count = 0
    with open('input.txt') as file:
        for line in file.readlines():
            brackets = re.findall(pattern, line)
            in_brackets = False
            for b in brackets:
                line = line.replace(b, '')
                in_brackets = in_brackets or is_palindrome(b)
            if not(in_brackets):
                line = line.split('[]')
                is_abba = False
                for l in line:
                    is_abba = is_abba or is_palindrome(l.strip())
                if (is_abba):
                    count = count + 1
    print('IPs :', count)


main()
