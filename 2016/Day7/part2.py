import re


def is_bab(string):
    # print('Checking ', string, ' for bab')
    bab = ''
    for i in range(0, len(string) - 2):
        word = string[i:i + 3]
        if (word[0] == word[2]) and (word[1] != word[0]):
            bab = word
            break
    # print('Result: ', bab)
    return bab


def is_aba(string, babs):
    aba = False
    abas = []
    # print('ABA: Checking ', string, ' having bab:', babs)
    for i in range(0, len(string) - 2):
        word = string[i:i + 3]
        if (word[0] == word[2]) and (word[1] != word[0]):
            bab = word[1] + word[0] + word[1]
            abas.append(bab)
            # print('ABA :', word, '  Check if ', bab, ' is in ', babs)
            if (bab in babs):
                aba = True
                break
    return aba


def main():
    pattern = r'(?<=\[).+?(?=\])'
    count = 0
    with open('input.txt') as file:
        for line in file.readlines():
            brackets = re.findall(pattern, line)
            babs = []
            # print('BAB : Checking: ', brackets)
            for b in brackets:
                line = line.replace(b, '')
                if (is_bab(b) != ''):
                    babs.append(is_bab(b))
            if len(babs) != 0:
                line = line.split('[]')
                abas = []
                is_abba = False
                for l in line:
                    if (is_bab(l) != ''):
                        aba = is_bab(l)
                        aba = aba[1] + aba[0] + aba[1]
                        abas.append(aba)
                if len((set(abas).intersection(set(babs)))) != 0:
                    count = count + 1
    print('IPs :', count)


main()
