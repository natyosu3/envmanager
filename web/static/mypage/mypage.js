const modal = document.getElementById('myModal');
    
function openModal() {
    modal.style.display = 'block';
}

function closeModal() {
    modal.style.display = 'none';
}

window.onclick = function(event) {
    if (event.target == modal) {
        closeModal();
    }
}


function addInput() {
    const inputContainer = document.getElementById('inputContainer');
    const inputGp = document.createElement('div');
    inputGp.classList.add('inputGp');
    inputGp.innerHTML = `
        <input type="text" placeholder="ENV_NAME" name="env_name">
        <input type="text" placeholder="ENV_VALUE" name="env_value">
        <button type="button" onclick="removeInput(this)">削除</button>
    `;
    inputContainer.appendChild(inputGp);
}


function removeInput(button) {
    const inputGp = button.parentNode;
    const inputContainer = inputGp.parentNode;
    inputContainer.removeChild(inputGp);
}