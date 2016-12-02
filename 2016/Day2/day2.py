# yup..this is lame
matrix = [[0] * 7 for i in range(7)]
matrix[1][3] = 1
matrix[2][2] = 2
matrix[2][3] = 3
matrix[2][4] = 4
matrix[3][1] = 5
matrix[3][2] = 6
matrix[3][3] = 7
matrix[3][4] = 8
matrix[3][5] = 9
matrix[4][2] = 'A'
matrix[4][3] = 'B'
matrix[4][4] = 'C'
matrix[5][3] = 'D'


def get_digit(cmd, start):
    digit = start
    for c in cmd:
        if c == 'L':
            start['j'] = start['j'] - 1 if (matrix[start['i']][(start['j'] - 1)]) != 0 else start['j']
        elif c == 'R':
             start['j'] = start['j'] + 1 if (matrix[start['i']][(start['j'] + 1)]) != 0 else start['j']
        elif c == 'U':
             start['i'] = start['i'] - 1 if (matrix[(start['i'] - 1)][start['j']]) != 0 else start['i']
        elif c == 'D':
             start['i'] = start['i'] + 1 if (matrix[(start['i'] + 1)][start['j']]) != 0 else start['i']
    return digit


def main():
    global matrix
    code = ''
    start = {'i': 3, 'j': 1}
    with open('input.txt') as file:
        input = file.readlines()
    for cmd in input:
        start = get_digit(cmd.strip(), start)
        code = code + str(matrix[start['i']][start['j']])
    print('Bathroom code: ', code)


main()
