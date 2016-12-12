def get_occurences(list):
    arr = [0] * 32
    list = ''.join(list)
    for c in list:
        if (arr[ord(c) - 97] == 0):
            arr[ord(c) - 97] = list.count(c)
    return arr


def get_max_characters(list, nr):
    i = 0
    arr = []
    while (i < nr):
        max_c = max(list)
        c = chr(list.index(max_c) + 97)
        arr.append(c)
        list[list.index(max_c)] = -1
        i = i + 1
    return arr


def main():
    sum = 0
    with open('input.txt') as input:
        for line in input.readlines():
            chars = line.strip().split('-')
            code = chars[-1].split('[')[0]
            sequence = chars[-1].split('[')[1][:-1]
            arr = get_occurences(chars[:-1])
            maxarr = get_max_characters(arr, 5)
            if (''.join(maxarr) == sequence):
                sum = sum + int(code)

    print('Sum: ', sum)


main()
