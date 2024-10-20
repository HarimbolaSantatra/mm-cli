from bs4 import BeautifulSoup
import os
import sys

# Search for a tag who has a class 'class_name'
def get_class(tag, class_name):
    return tag.find_all("span", class_=class_name)

def main():
    html_filename = "result.html"

    # get path of the executable
    file_path = os.path.dirname(os.path.abspath(__file__))

    # open, read and close file
    fp = open(os.path.join(file_path, html_filename), "r")
    soup = BeautifulSoup(fp, 'html.parser')
    fp.close()

    entry_word = get_class(soup, "entryWord")
    print(entry_word)

if __name__ == '__main__':
    main()
