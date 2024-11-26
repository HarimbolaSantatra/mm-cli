# Test the HTML scraper
# Test if the scraper needs an update.
# That's the case only if the website has been updated by the admin
# and it's not in the same structure anymore.

import unittest
import subprocess
import json
import sys

HTML_TESTFILE = "scraping-test/result.html"
JSON_RESULT = "scraping-test/test-result.json"

class TestScraping(unittest.TestCase):

    def test_scraper(self):

        # Create and run commands
        # Full command is "cat scraping-test/result.html | xargs -0 ./mm-scraping"
        cat_process = subprocess.Popen(["cat", HTML_TESTFILE],
                                       stdout=subprocess.PIPE,
                                       text=True)
        xarg_process = subprocess.Popen(["xargs", "-0", "./mm-scraping"],
                                        stdin=cat_process.stdout,
                                        stdout=subprocess.DEVNULL,
                                        stderr=None,
                                        text=True)

        try:
            output, error = xarg_process.communicate(timeout=15)
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
        self.assertEqual(jsfile["Partie du discours\u00a0"], "nom")
        self.assertEqual(jsfile["Vocabulaire\u00a0"], "Economie: alimentation")
        self.assertEqual(jsfile["Explications en malgache\u00a0"], "Trano fanaovana nahandro:Tsy miala ao an-dakozia foana io saka io[1.1]")
        self.assertEqual(jsfile["Explications en fran\u00e7ais\u00a0"], "cuisine[1.3]")

if __name__ == '__main__':
    unittest.main()
