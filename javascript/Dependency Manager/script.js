var inputArrf = [ // it is an array as asked in the excercise and it would've been much easier with a JSON
    "G: D",
    "D:  ",
    "A: G",
    "E: A",
    "B: C",
    "C: D",
    "F: B",
    "K: D"
    
]

function InstallationManager(inputArr) {
// this code is to transfer the array to a JSON Object
var myObject = {}

for(i = 0; i < inputArr.length; i++) {
    let holder = inputArr[i].replace(/\s/g, ''); //removing spaces
    holder = holder.split(":"); // getting an array of [key, value]
    myObject[holder[0]] = holder[1]; // making up the Object    
}

var MyKeysArr = Object.keys(myObject);
var MyValsArr = Object.values(myObject);
var outputArr = []; // if something makes it to this array then it must've been installed
var indexHolder = [];

// pushing non-dependent packages. There should at least be one in all times
for (j = 0; j < MyValsArr.length; j++){
    if (MyValsArr[j].trim() == "") {
        outputArr.push(MyKeysArr[j])
    } 
}
// grabbing the indexes of all of the packages without dependencies that are actually used as dependecies later
    for (k = 0; k < outputArr.length; k++){
        if (MyValsArr.includes(outputArr[k])) {
            indexHolder[k] = k; 
        }
    } 

    for (m = 0; m < MyValsArr.length; m++){
        let Installed = outputArr[m];
        let indexOfNextDep = MyValsArr.indexOf(Installed)
        let nextPack = MyKeysArr[MyValsArr.indexOf(Installed)];
        while (indexOfNextDep != -1 && nextPack){       // this while loop fixes the problem of multiple indeces for a given dependency
           outputArr.push(MyKeysArr[indexOfNextDep]);   // this page saved the day for me: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/indexOf
           indexOfNextDep = MyValsArr.indexOf(Installed, indexOfNextDep + 1);
        }
        // if (nextPack){                               // this is how it used to be but with the above while loop fixes the issue.
        // outputArr.push(nextPack);
        // } 
    }

    // code to that will display rejection in the case of a dependency cycle
    var DepsAndPacks = MyKeysArr.concat(MyValsArr);
    DepsAndPacks = DepsAndPacks.filter(entry => /\S/.test(entry)); // had to look through MDN to know what .test() function does.
    unique = a => [...new Set(a)]; /* couldn't resist ES6 Hottness and totally stackoverflowed this one.
                                      sorry this was the only thing that wasn't my idea */
    if (unique(DepsAndPacks).length <= outputArr.length){
        return outputArr.toString();
    } else {
       throw 'There is a Cycle in the package that you entered'
    }
}


// PLEASE READ THESE NOTES: 
/*  this solution works with arrays weather they have cycle in them or not because it installs everything
outside of the cycle and ignores the packages within the cycle, but if you want to throw an error message
then please feel free to un-comment the code between the lines: 53 -> 61. */