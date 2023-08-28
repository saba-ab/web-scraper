import puppeteer from "puppeteer";
const url = "https://www.ambebi.ge/";
const keyWord = "საბა";
async function main() {
  const browser = await puppeteer.launch();
  const page = await browser.newPage();
  await page.goto(url);
  const linksWithText = await page.$$eval("a", (anchors) =>
    anchors.map((anchor) => {
      const nestedP = anchor.querySelector("p");
      return {
        anchor: anchor.href,
        nestedP: nestedP
          ? nestedP.textContent.trim()
          : "there is no nested text",
      };
    })
  );
  let newsArr = [];
  linksWithText.forEach(({ anchor, nestedP }) => {
    if (nestedP !== "there is no nested text") {
      const news = {
        link: anchor,
        Text: nestedP,
      };
      newsArr = [...newsArr, ...[news]];
    }
  });
  console.log(newsArr.length);
  newsArr.forEach((news) => {
    if (news.Text.includes(keyWord)) {
      console.log(news);
    }
  });

  await browser.close();
}
main();
