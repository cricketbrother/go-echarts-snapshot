'use strict';

const argv = process.argv;
let puppeteerPath, browserPath, renderer, src, dst, width, height;
let setViewportFlag = false;
if (argv.length === 7) {
    puppeteerPath = argv[2];
    browserPath = argv[3];
    renderer = argv[4];
    src = argv[5];
    dst = argv[6];
} else if (argv.length === 9) {
    puppeteerPath = argv[2];
    browserPath = argv[3];
    renderer = argv[4];
    src = argv[5];
    dst = argv[6];
    width = argv[7];
    height = argv[8];
    setViewportFlag = true;
} else {
    console.log('Usage: \n' +
        '> node screenshot.js puppeteerPath browserPath renderer src dst' +
        '> node screenshot.js puppeteerPath browserPath renderer src dst width height');
    process.exit(1);
}

const puppeteer = require(puppeteerPath);


(async () => {
    const browser = await puppeteer.launch({
        executablePath: browserPath,
        args: ['--no-sandbox']
    });
    const page = await browser.newPage();
    if (setViewportFlag) {
        await page.setViewport({width: width, height: height});
    }
    await page.goto(src);
    const element = await page.$(renderer);
    await element.screenshot({path: dst})
    await browser.close();
})();
