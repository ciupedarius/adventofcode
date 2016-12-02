def get_digit(cmd, start):
    digit = start
    for c in cmd:
        if c == 'U':
            digit = digit - 3 if digit - 3 > 0 else digit
        elif c == 'D':
            digit = digit + 3 if digit + 3 < 10 else digit
        elif c == 'L':
            digit = digit - 1 if (digit - 1) % 3 != 0 else digit
        elif c == 'R':
            digit = digit + 1 if digit % 3 != 0 else digit
    return digit


def main():
    code = ''
    start = 5
    with open('input.txt') as file:
        input = file.readlines()
    for cmd in input:
        start = get_digit(cmd.strip(), start)
        code = code + str(start)
    print('Bathroom code: ', code)


main()
