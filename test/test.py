# Test the HTML parser script (`mm-parsing`)
# Test if the scraper needs an update.
# That's the case only if the website has been updated by the admin
# and it does not have the same HTML layout anymore.

import unittest
import subprocess
import json
import sys

HTML_TESTFILE = "scraping-test/result.html"
JSON_RESULT = "scraping-test/parsing-result.json"

class TestScraping(unittest.TestCase):

    def test_scraper(self):

        html_file = open(HTML_TESTFILE, "r")
        xarg_process = subprocess.Popen(["./mm-parsing"], stdin=html_file, text=True,
                                        stdout=subprocess.PIPE)
        html_file.close()

        try:
            output, _ = xarg_process.communicate(timeout=15)
            print(output)
            xarg_process.kill()

        except subprocess.TimeoutExpired:
            xarg_process.kill()
            print(f"Timeout error")
            sys.exit(3)

        # Open JSON file into a variable
        with open(JSON_RESULT, 'r') as f:
            jsfile = json.load(f)

        # Assert each field of the json file
        self.assertEqual(jsfile["Part of speech\u00a0"], "2noun")
        self.assertEqual(jsfile["Vocabulary\u00a0"], "5Economy: food")
        self.assertEqual(jsfile["Explanations in Malagasy\u00a0"], "3[1.1]Trano fanaovana nahandro:Tsy miala ao an-dakozia foana io saka io")
        self.assertEqual(jsfile["Explanations in French\u00a0"], "4[1.3][lakoZY] cuisine")

if __name__ == '__main__':
    unittest.main()
