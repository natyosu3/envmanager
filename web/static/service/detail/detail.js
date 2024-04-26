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

function decrypt() {
    var encrypted = document.getElementById("encrypted-env").value;
    var password = document.getElementById("password").value;
    var decrypted = CryptoJS.AES.decrypt(encrypted, password);
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
    const serviceModel = createServiceModel(jsonData);

    document.getElementById("envs").value = JSON.stringify(serviceModel.envs);

    deleteView();
    createEnvView(serviceModel);
}

function deleteView() {
    var preContentElement = document.querySelector('.pre-content');
    if (preContentElement) {
        preContentElement.parentNode.removeChild(preContentElement);
    }
}

function createEnvView(env_data) {
    var target = document.getElementById("content");
    target.innerHTML = `<h2>${env_data.service_name}</h2>`;
    env_data.envs.forEach(env => {
        target.innerHTML += `
            <div class="env">
                <input type="text" name="env_name" value="${env.env_name}" readonly>
                <input type="text" name="env_value" value="${env.env_value}" readonly>
            </div>
        `
    });
    target.innerHTML += `
        <button type="button" onclick="createDotEnv()">.ENVファイルを作成</button>
        <button type="button" onclick="location.href='/service/edit/${env_data.service_id}'">編集</button>
        <button type="button" onclick="location.href='/service/dashboard'">戻る</button>
    `
}

function createDotEnv() {
    var envs = JSON.parse(document.getElementById('envs').value);
    var dotEnv = '';
    envs.forEach(env => {
        dotEnv += `${env.env_name}=${env.env_value}\n`;
    });

    var blob = new Blob([dotEnv], { type: 'text/plain' });
    var url = window.URL.createObjectURL(blob);
    var a = document.createElement('a');
    alert("ダウンロードを実行します. ファイル名を.envとして保存してください.")
    a.download = ".env";
    a.href = url;
    a.click();
    window.URL.revokeObjectURL(url);
}