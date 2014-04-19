/*
 * easy problem that average student can solve quickly.
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
    var testNum = args.testCaseNum;
    var inputs = args.inputs;
    for (var i = 0; i < testNum; i++) {
        var height = 1;
        for (var j = 0; j < inputs[i]; j++) {
            height = j % 2 == 0 ? height * 2 : height + 1;
        }
        process.stdout.write("" + height + "\n");
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
    var numericInputs = [];
    lines.splice(0, 1);
    lines.forEach( function (num) {
        numericInputs.push(parseInt(num));
    });
    return { 'testCaseNum' : testCaseNum, 'inputs' : numericInputs }
}

/*
 * put print out debug code into here
 * call debug() in solver to print 
 */
function debug() {
    
}

/*================ helpers for this specific problem  =================*/

