reg = [0] * 4

# uncomment this for part2
# reg[2] = 1


def pos(arg):
    return ord(arg) - 97


def doOp(cmd, arg1, arg2):
    global reg
    if cmd == 'cpy':
        if arg1.isdigit():
            reg[pos(arg2)] = int(arg1)
        else:
            reg[pos(arg2)] = reg[pos(arg1)]
    elif cmd == 'inc':
        reg[pos(arg1)] = reg[pos(arg1)] + 1
    elif cmd == 'dec':
        reg[pos(arg1)] = reg[pos(arg1)] - 1
    elif cmd == 'jnz':
        if arg1.isdigit() and arg1 != 0:
            return int(arg2)
        elif not(arg1.isdigit()) and reg[pos(arg1)] != 0:
            return int(arg2)
        else:
            return 1
    return 1


def main():
    global reg
    cmds = []
    with open('input.txt') as input:
        for line in input.readlines():
            cmds.append(line.strip())
    i = 0
    while (i < len(cmds)):
        cmd = cmds[i].split()
        inc = 0
        if len(cmd) == 2:
            inc = doOp(cmd[0], cmd[1], None)
        elif len(cmd) == 3:
            inc = doOp(cmd[0], cmd[1], cmd[2])
        i = i + inc
    print('Register a:', reg[0])


main()
