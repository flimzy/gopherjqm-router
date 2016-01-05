if (typeof(window) === 'undefined') {
    // This is necessary because gopherjs runs the test from /tmp
    // rather than from the current directory, which confuses nodejs
    // as to where to search for modules
    var cwd = process.cwd();
    var jsdom = require(cwd+'/node_modules/jsdom');
    var doc = jsdom.jsdom("<html></html>");
    global.window = doc.defaultView;
    global.jQuery = require(cwd + '/node_modules/jquery');
}
