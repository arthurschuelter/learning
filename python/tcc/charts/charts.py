import numpy as np

def count_links(links):
    d = dict()

    for l in links:
        nl = l.replace('https://', '')
        idx = nl.find('/')
        nl = nl[:idx]

        if not nl in d:
            d[nl] = 1
        else:
            d[nl] = d[nl] + 1

    return d

filename = 'links.in'
f = open(filename, 'r')

links = f.readlines()
d = count_links(links)
for key in d:
    print(key, '->', d[key])

f.close()