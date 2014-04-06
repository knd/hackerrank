/*
 * I go with instinct:
 * - can i do with linear time? i coulnd't think of a way so i aim for N(logN)
 * - i sort the array first, things i learn:
 *      + just use the built-in array sort, i wasted time writing my own quicksort and had to deal with bunch of NaN, undefined of JS array. 
 *      + array built-in sort uses stable merge sort algorithm, switching from heap sort in 2004. read more: https://bugzilla.mozilla.org/show_bug.cgi?id=224128
 *      + this sort method sorts lexicographically according to string conversion, so i had to pass in my own compare function
 * - with sorted array, iterate it and grab K=3 elements (eg. from index 3 to index 5) and get difference between the values of those 2 indexes
 * - find the min difference after iterating the whole array
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
    var totalPacket = args.totalPacket;
    var disPacket = args.disPacket;
    var lstOfCandyCount = args.lstOfCandyCount;

    lstOfCandyCount.sort(function (a, b) { return a - b; });
    if (disPacket <= 1) {
        process.stdout.write('0');
    } else {
        var minDiff = 1000000000; // 10^9
        var i = disPacket - 1;
        while (i < totalPacket) {
            var diff = lstOfCandyCount[i] - lstOfCandyCount[i-disPacket+1];
            if (diff < minDiff) {
                minDiff = diff;
            }
            i++;
        }
        process.stdout.write(minDiff.toString());
    }
    
}


/*================ generic helpers for most problems =================*/

/*
 * process all the raw input from command line
 * should modify with care
 */
function processInput(input) {
    var lines = input.split('\n');
    var totalPacket = parseInt(lines[0]);
    var disPacket = parseInt(lines[1]);
    var lstOfCandyCount = [];
    lines.splice(0, 1);
    lines.splice(0, 1);
    lines.forEach( function (num) {
        lstOfCandyCount.push(parseInt(num));
    });
    return { 
        'totalPacket' : totalPacket, 
        'disPacket' : disPacket,
        'lstOfCandyCount' : lstOfCandyCount
    };
}

/*
 * put print out debug code into here
 * call debug() in solver to print 
 */
function debug(whatToPrint, exit) {
    console.log(whatToPrint);
    if (exit) {
        process.exit();
    }
}

/*================ helpers for this specific problem  =================*/

