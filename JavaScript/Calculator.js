const num1 = document.getElementById('num1');
const num2 = document.getElementById('num2');
const buttonadd = document.getElementById('add-button');
const answer = document.getElementById('answer');

buttonadd.addEventListener('click', () => {
    sum = parseInt(num1.value) + parseInt(num2.value);
    answer.textContent = `The answer is ${sum}`;
})