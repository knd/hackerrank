/*
 * This is an easy problem which makes some brain blood flowing.
 */

process.stdin.resume();
process.stdin.setEncoding('ascii');

/* 
 * initialize globale vars 
 */
var input = '';

/* 
 * read inputs from the command line 
 */
process.stdin.on('data', function (data) {
    input += data;
});

/* 
 * finish reading inputs from command line 
 * 1. process inputs for necessary correct types and values 
 * 2. pass necessary inputs into solver params to solve 
 */
process.stdin.on('end', function () {
    // probably should modify processInput with care
    var processed = processInput(input);        
    solver(processed);                          
});

/* 
 * actual code with standard output go into here to solve problem 
 * 'args' is a javascript dict {} 
 */
function solver(args) {
    // YOUR MAIN CODE STARTS HERE
    var testCaseNum = args.testCaseNum;
    var inputs = args.inputs;
    var suite, money, cost, promo;
    var chocolateLeft, chocolateEat;
    for (var i = 0; i < testCaseNum; i++) {
        suite = inputs[i];
        money = suite[0];
        cost = suite[1];
        promo = suite[2];
        
        chocolateLeft = Math.floor(money / cost);
        chocolateEat = 0;
        while (chocolateLeft >= promo) {
            chocolateEat += promo;
            chocolateLeft = chocolateLeft - promo + 1;
        }
        if (chocolateLeft < promo) {
            chocolateEat += chocolateLeft;
        }
        process.stdout.write(chocolateEat + '\n');
    }
}


/*================ generic helpers for most problems =================*/

/*
 * process all the raw input from command line
 * should modify with care
 */
function processInput(input) {
    var lines = input.split('\n');
    var testCaseNum = parseInt(lines[0]);
    var inputSuite = [];
    lines.splice(0, 1);
    lines.forEach( function (suite) {
        var arr = suite.split(' ').map( function (item) {
            return parseInt(item); 
        });
        inputSuite.push(arr);
    });
    return { 'testCaseNum' : testCaseNum, 'inputs' : inputSuite }
}

/*
 * put print out debug code into here
 * call debug() in solver to print 
 */
function debug() {
    
}

/*================ helpers for this specific problem  =================*/

