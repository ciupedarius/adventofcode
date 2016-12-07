def is_triangle(corners):
    corners.sort()
    if corners[0] + corners[1] > corners[2]:
        return True
    else:
        return False


def main():
    no_tri = 0
    with open('input.txt') as file:
        input = file.readlines()
    for tri in input:
        if is_triangle(list(map(int, tri.strip().split()))):
            no_tri += 1
    print('Triangles: ', no_tri)

main()
