const modal = document.getElementById('myModal');

function openModal() {
    modal.style.display = 'block';
}

function closeModal() {
    modal.style.display = 'none';
}

window.onclick = function (event) {
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



class Env_model {
    constructor(env_name, env_value) {
        this.env_name = env_name;
        this.env_value = env_value;
    }
}

class Service_model {
    constructor(service_name, service_id, envs) {
        this.service_name = service_name;
        this.service_id = service_id;
        this.envs = envs;
    }
}

function createRandomId(length) {
    var result = '';
    var characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    var charactersLength = characters.length;
    for (var i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
}


function convertJson() {
    if (document.getElementById('encrypt_key').value == '') {
        alert('暗号化キーを入力してください');
        return
    }

    if (!checkEncryptKey()) {
        return
    }

    var env_names = document.getElementsByName('env_name');
    var env_values = document.getElementsByName('env_value');

    var envs = [];
    for (var i = 0; i < env_names.length; i++) {
        envs.push(new Env_model(env_names[i].value, env_values[i].value));
    }

    var service_name = document.getElementById('service_name').value;
    var service_id = createRandomId(12);
    document.getElementById('service_id').value = service_id;

    var service = new Service_model(service_name, service_id, envs);
    var json = JSON.stringify(service);

    var encrypt_key = document.getElementById('encrypt_key').value;
    if (encrypt_key != '') {
        var encrypted = CryptoJS.AES.encrypt(json, encrypt_key).toString();
        json = encrypted;
    }

    document.getElementById('json').value = json;
}

function checkEncryptKey() {
    var encrypt_key = document.getElementById('encrypt_key').value;
    var confirm_encrypt_key = document.getElementById('confirm_encrypt_key').value;

    if (encrypt_key != confirm_encrypt_key) {
        alert('暗号化キーが一致しません');
        return false;
    }
    return true;
}



document.getElementById("envForm").addEventListener("submit", function (event) {
    event.preventDefault(); // フォームのデフォルトの送信をキャンセル

    var jsonarea = document.getElementById('json');
    if (jsonarea.value == '') {
        alert('JSONを生成してください');
        return
    }
    var password = document.getElementById('encrypt_key').value;
    if (password == '') {
        alert('パスワードを入力してください');
        return
    }

    // POSTリクエストを送信
    fetch('http://localhost:8080/service/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            service_id: document.getElementById('service_id').value,
            service_name: document.getElementById('service_name').value,
            data: jsonarea.value
        })
    })
        .then(data => {
            window.location.href = data.url;
        })
        .catch(error => {
            console.error('There was a problem with the POST request:', error);
        });
});