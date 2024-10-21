import bs4
import os
import sys
import typing

# Search for a tag who has a class 'class_name'
def get_class(tag, class_name) -> bs4.Tag :
    return tag.find("span", class_=class_name)

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
        soup = bs4.BeautifulSoup(fp, 'html.parser')
        fp.close()

        # self.soup is a <html> tag object
        self.soup = soup

    def get_entry_word(self) -> str:
        """
        Get the keywords entered by the user
        In french: Entrée
        """
        # find the 'entryWord' class
        entry_word_tag: bs4.Tag = get_class(self.soup, "entryWord")
        # the result of the above expression is always in the format:
        # <span class="entryWord">lako<u>zi</u>a<a class="entryWord" name="mg.n"> </a></span>
        # ... so we need to flatten it
        entry_word = ""
        for child in entry_word_tag.children:
            # append to result if 'child' is a NavigableString
            entry_word += child if type(child) == bs4.element.NavigableString else child.string
        return entry_word


    def get_tables(self) -> list:
        """
        Get all the tables in the <body>
        We need to get all the tables because each line in the result is represented by a <table>

        Returns
        -------
        list
            Each row of the list has the format:
                {
                    "label": "Explication en malgache",
                    "text": "Trano fanaovana nahandro"
                }
        """
        tables: list = self.soup.find_all("table")
        result = []

        for table in tables[4:10]:
            # table[4:] because we do not want to print the first 4 <table>,
            # which is the navigation menu and some unknown BS!
            # TODO: this needs a hard edit cuz that guy was fkn lazy!

            new_entry = {}
            # take a look at the html file (result.html) to understand this section
            row = table.find("tr")
            new_entry["label"] = row.contents[0].stripped_strings
            new_entry["text"]  = "".join(row.contents[1].stripped_strings)
            result.append(new_entry)

        return result


def main():
    html_filename = "result.html"

    # get path of the executable
    file_path = os.path.dirname(os.path.abspath(__file__))
    filename = os.path.join(file_path, html_filename)

    doc = Document(filename)
    tables = doc.get_tables()
    inc = 1
    for table in tables:
        print("{}- {}: {}".format(inc, table["label"], table["text"]))
        inc += 1


if __name__ == '__main__':
    main()
