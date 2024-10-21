from bs4 import BeautifulSoup
import os
import sys

# Search for a tag who has a class 'class_name'
def get_class(tag, class_name):
    return tag.find_all("span", class_=class_name)

class Document:
    """
    Represent a HTML document in response to a query from motmalagasy.org
    """

    def __init__(self, filename):
        """
        Constructor

        Parameters
        ----------
        filename: string
            get the name of the HTML file
        """
        # open, read and close file
        fp = open(filename, "r")
        soup = BeautifulSoup(fp, 'html.parser')
        fp.close()

        # self.soup is a <html> tag object
        self.soup = soup

    def get_entry_word(self):
        """
        Get the keywords entered by the user
        """
        # find the 'entryWord' class
        entry_word = get_class(self.soup, "entryWord")
        return entry_word


def main():
    html_filename = "result.html"

    # get path of the executable
    file_path = os.path.dirname(os.path.abspath(__file__))
    filename = os.path.join(file_path, html_filename)

    doc = Document(filename)
    print(doc.get_entry_word())


if __name__ == '__main__':
    main()
