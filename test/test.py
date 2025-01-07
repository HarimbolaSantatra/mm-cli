# Test the HTML parser script (`mm-parsing`)
# Test if the scraper needs an update.
# That's the case only if the website has been updated by the admin
# and it does not have the same HTML layout anymore.

import unittest
import subprocess
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
            xarg_process.kill()

        except subprocess.TimeoutExpired:
            xarg_process.kill()
            print(f"Timeout error")
            sys.exit(3)

        # Open JSON file into a variable
        with open(JSON_RESULT, 'r') as f:
            jsfile = f.read()

        self.assertEqual(jsfile, output)

if __name__ == '__main__':
    unittest.main()
