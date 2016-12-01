directions = []   # input directions
orientation = 1  # north first
visited = []    # visited coords
target = {}     # target coords


def move_units(coords, arg, units, op):

    global visited
    global target
    for i in range(0, units):
        coords[arg] = coords[arg] + 1 if op == 0 else coords[arg] - 1
        obj = {'x': coords['x'], 'y:': coords['y']}
        if not(obj in visited):
            visited.append({'x': coords['x'], 'y:': coords['y']})
        else:   # found first duplicate, save coords
            if (target == {}):
                target = {'x': coords['x'], 'y': coords['y']}
            break
    return coords[arg]


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
        coords['y'] = move_units(coords, 'y', units, 0)
    elif (orientation == 2):  # west
        coords['x'] = move_units(coords, 'x', units, 1)
    if (orientation == 3):  # south
        coords['y'] = move_units(coords, 'y', units, 1)
    elif (orientation == 0):  # east
        coords['x'] = move_units(coords, 'x', units, 0)
    return coords


def main():
    global visited
    global target
    with open('input.txt') as input:
        line = input.readline()
        directions = line.split(', ')
    coords = {'x': 0, 'y': 0}
    # add origin to visited
    visited.append({'x': coords['x'], 'y:': coords['y']})
    for dir in directions:
        coords = move(dir[0], dir[1:], coords)
    print("Blocks: ", abs(target['x']) + abs(target['y']))


main()
