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

        # Create and run commands
        # Full command is "cat scraping-test/result.html | xargs -0 ./mm-parsing"
        cat_process = subprocess.Popen(["cat", HTML_TESTFILE],
                                       stdout=subprocess.PIPE,
                                       text=True)
        xarg_process = subprocess.Popen(["xargs", "-0", "./mm-parsing"],
                                        stdin=cat_process.stdout,
                                        stdout=subprocess.DEVNULL,
                                        stderr=None,
                                        text=True)

        try:
            output, error = xarg_process.communicate(timeout=15)
            if error != None:
                print(error, file=sys.stderr)
                sys.exit(1)

            xarg_process.kill()
        except TimeoutExpired:
            cat_process.stdout.close()
            cat_process.kill()
            xarg_process.kill()
            print("Timeout for xargs process")
            sys.exit(1)

        cat_process.stdout.close()
        cat_process.kill()


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
