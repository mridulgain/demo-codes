// Printing PI value
// console.log(Math.PI);  // Prints: 3.141592653589793

// Function to calculate circle area
function calculateCircleArea(radius){
    var area = (Math.PI) * radius * radius;
    return area;
}
 
// console.log(calculateCircleArea(5));  // Prints: 78.53981633974483

// console.log(calculateCircleArea(10));
// console.log(Math.round(16777215).toString(16))
//console.log("<br>")
var x = [1,2,3,45,6]
for(num of x){
    console.log(num)
}

for(i in x){
    console.log(i)
}

function randomColor(){
    return Math.floor(Math.random()*256);
}

function setColor(){
    r = randomColor();
    g = randomColor();
    b = randomColor();
    c = "rgb(" + r + "," + g + "," + b + ")"
    document.body.style.backgroundColor = c;
    color_code.innerHTML = c;
}

function any(){
    n = document.getElementById("number").value;
    console.log(n * 2)
}

function change_col(){
    c = "rgb(" + red.value + "," + green.value + "," + blue.value + ")"
    document.body.style.backgroundColor = c;
}

// document.body.style.backgroundImage = "url(day.jpg)"
// document.getElementById("test").style.backgroundColor = red