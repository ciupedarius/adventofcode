directions = []  # input directions
orientation = 1  # start facing north


def move(direction, units, coords):
    global orientation
    units = int(units)
    if (direction == 'L'):
        orientation += 1
        orientation = orientation % 4
    elif (direction == 'R'):
        orientation -= 1
        orientation = orientation % 4
    if (orientation == 1):  # north
        coords['y'] = coords['y'] + units
    elif (orientation == 2):  # west
        coords['x'] = coords['x'] - units
    if (orientation == 3):  # south
        coords['y'] = coords['y'] - units
    elif (orientation == 0):  # east
        coords['x'] = coords['x'] + units
    return coords


def main():
    with open('input.txt') as input:
        line = input.readline()
        directions = line.split(', ')
    coords = {'x': 0, 'y': 0}
    for dir in directions:
        coords = move(dir[0], dir[1:], coords)
    print("Blocks: ", abs(coords['x']) + abs(coords['y']))


main()
