def is_triangle(corners):
    corners.sort()
    if corners[0] + corners[1] > corners[2]:
        return True
    else:
        return False


def main():
    no_tri = 0
    i = 0
    col1 = []
    col2 = []
    col3 = []
    with open('input.txt') as file:
        for line in file:
            columns = line.split()
            col1.append(columns[0])
            col2.append(columns[1])
            col3.append(columns[2])
            i += 1
            if i == 3:
                if is_triangle(list(map(int, col1))):
                    no_tri += 1
                if is_triangle(list(map(int, col2))):
                    no_tri += 1
                if is_triangle(list(map(int, col3))):
                    no_tri += 1
                i = 0
                col1 = []
                col2 = []
                col3 = []
    print('Triangles: ', no_tri)   


main()
