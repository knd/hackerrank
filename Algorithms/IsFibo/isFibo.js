/*
 * To solve this problem, I go with instinct
 * - need a fibonacci sequence to check whether a number is in it
 * - need to generate a fibonacci sequence in advance before check
 * - need to only generate fibonacci sequence up to somewhere the number being checked
 * - need to keep fibonacci sequence live for further test case check or caching if like to call
 * - need to use binary search in fibonacci sequence for log(N) speed in case there is time constraint
 */

process.stdin.resume();
process.stdin.setEncoding('ascii');

/* 
 * initialize globale vars 
 */
var input = '';
var fibonacciSeq = [0, 1];

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
    var testCaseNum = args.testCaseNum;
    var inputs = args.inputs;
    for (var i = 0; i < testCaseNum; i++) {
        var num = inputs[i];
        if (isFibo(num)) {
            process.stdout.write('IsFibo\n');
        } else {
            process.stdout.write('IsNotFibo\n');
        }
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

function isFibo(num) {
    var len = fibonacciSeq.length;
    if (fibonacciSeq[len-1] < num) {
        var curNum = fibonacciSeq[len-1];
        var i = len-1;
        var newNum;
        while (curNum < num) {
            newNum = fibonacciSeq[i] + fibonacciSeq[i-1];
            fibonacciSeq.push(newNum);
            i++;
            curNum = fibonacciSeq[i];
            if (curNum == num) {
                return true;
            }
        }
        return false;
    } else {
        return contains(num, fibonacciSeq, 0, len-1);
    }
}

function contains(num, fibonacciSeq, low, high) {
    while (low <= high) {
        var mid = Math.floor((low + high) / 2);
        if (fibonacciSeq[mid] == num) {
            return true;
        } else if (fibonacciSeq[mid] < num) {
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }
    return false;
}
