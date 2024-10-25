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


    def get_result_pairs(self) -> list:
        """
        Get all the pair of result.
        e.g:
            Entrée: lakozia
            Partie du discours: nom
            Vocabulaire Economie: alimentation

        We need to get a table tag, which some particular attribute:
            - cellspacing="4"
            - width="100%"

        Returns
        -------
        list
            Each row of the list has the format:
                {
                    "label": "Explication en malgache",
                    "text": "Trano fanaovana nahandro"
                }
        """
        tables: list = self.soup.find_all(name="td", width="20%", align="right")
        result = []

        for table in tables:
            # table[4:] because we do not want to print the first 4 <table>,
            # which is the navigation menu and some unknown BS!
            # TODO: this needs a hard edit cuz that guy was fkn lazy!

            new_entry = {}
            # take a look at the html file (result.html) to understand this section
            new_entry["label"] = table.string

            # put next_sibling 2 times because look at https://www.crummy.com/software/BeautifulSoup/bs4/doc/#next-sibling-and-previous-sibling
            new_entry["text"]  = "".join(table.next_sibling.next_sibling.stripped_strings)

            result.append(new_entry)

        return result


def main():
    html_filename = "result.html"

    # get path of the executable
    file_path = os.path.dirname(os.path.abspath(__file__))
    filename = os.path.join(file_path, html_filename)

    doc = Document(filename)
    tables = doc.get_result_pairs()
    inc = 1
    for table in tables:
        print("{}- {}: {}".format(inc, table["label"], table["text"]))
        inc += 1


if __name__ == '__main__':
    main()
