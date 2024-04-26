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



document.addEventListener('DOMContentLoaded', function () {
    service_name = document.getElementById('service_name');
    key = prompt(`サービス: ${service_name.value}の復号キーを入力してください`, "");

    if (key == null) {
        return;
    }

    encrypted = document.getElementById('json').value;
    serviceModel = decrypt(encrypted, key);
    createEnvView(serviceModel);
});


function decrypt(encrypted, key) {
    var decrypted = CryptoJS.AES.decrypt(encrypted, key);
    if (decrypted.toString(CryptoJS.enc.Utf8) == '') {
        alert('復号キーが正しくありません');
        return;
    }

    const jsonData = JSON.parse(decrypted.toString(CryptoJS.enc.Utf8));

    function createEnvModel(envData) {
        return new Env_model(envData.env_name, envData.env_value);
    }
    function createServiceModel(jsonData) {
        const envModels = jsonData.envs.map(envData => createEnvModel(envData));
        return new Service_model(jsonData.service_name, jsonData.service_id, envModels);
    }
    return createServiceModel(jsonData);
}

function createEnvView(env_data) {
    var target = document.getElementById("content");
    var bottom = document.getElementById("bottom");
    target.innerHTML = `<h2>${env_data.service_name}</h2>`;
    env_data.envs.forEach(env => {
        target.innerHTML += `
            <div class="env">
                <input type="text" name="env_name" value="${env.env_name}">
                <span>=</span>
                <input type="text" name="env_value" value="${env.env_value}">
                <button type="button" onclick="removeInput(this)">削除</button>
            </div>
        `
    });
    bottom.innerHTML += `
        <button type="button" onclick="addInput()">環境変数を追加</button>
        <input type="password" id="encrypt_key" placeholder="暗号化キーを入力(必須)" required>
        <input type="password" id="confirm_encrypt_key" placeholder="再確認" required>
        <button type="button" onclick="convertJson()">変更を確定</button>
        <button type="button" onclick="location.href='/service/dashboard'">戻る</button>
    `
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
    var service_id = document.getElementById('service_id').value;

    var service = new Service_model(service_name, service_id, envs);
    var json = JSON.stringify(service);

    var encrypt_key = document.getElementById('encrypt_key').value;
    if (encrypt_key != '') {
        var encrypted = CryptoJS.AES.encrypt(json, encrypt_key).toString();
        json = encrypted;
    }

    document.getElementById('json').value = json;
    createPostRequest();
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

function removeInput(button) {
    const inputGp = button.parentNode;
    const inputContainer = inputGp.parentNode;
    inputContainer.removeChild(inputGp);
}

function addInput() {
    const inputContainer = document.getElementById('content');
    const inputGp = document.createElement('div');
    inputGp.classList.add('inputGp');
    inputGp.innerHTML = `
        <input type="text" placeholder="ENV_NAME" name="env_name">
        <span>=</span>
        <input type="text" placeholder="ENV_VALUE" name="env_value">
        <button type="button" onclick="removeInput(this)">削除</button>
    `;
    inputContainer.appendChild(inputGp);
}

function createPostRequest() {
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
    fetch('http://localhost:8080/service/update', {
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
};