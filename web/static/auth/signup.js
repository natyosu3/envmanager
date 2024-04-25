function confirmPassword() {
    const password = document.getElementById('password').value;
    const confirmPassword = document.getElementById('confirm_password').value;
    const errorMsg = document.getElementById('error_msg');

    if (password == confirmPassword) {
        errorMsg.innerText = "";
        return true;
    } else {
        errorMsg.innerText = "パスワードが一致しません";
        return false;
    }
}