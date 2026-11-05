import requests

def remove_tags(string):
    stack = []
    aux = ""
    for i in range(0, len(string)):
        if( string[i] == '<' ):
            for j in range(i, len(string)):
                aux = aux + string[j]
                if( string[j] == '>'):
                    break
            
            stack.append( str(aux) )
            aux = ""

    for el in stack:
        string = string.replace(el,'')

    return string


def extract_publication(url):
    publication = dict()

    r = requests.get(url)
    code = r.text.split('\n')

    title = ""

    code = list( filter(lambda a: a != '', code) )
    for l in code:
        if( "<title>" in l ):
            title = remove_tags(l)
            title = title.split( ' | ' )[0]

    publication['title'] = title

    return publication

def main():

    publication = dict()
    url = "https://dl.acm.org/doi/proceedings/10.1145/502512"

    publication = extract_publication(url)

    print( publication )


if __name__ == "__main__":
    main()