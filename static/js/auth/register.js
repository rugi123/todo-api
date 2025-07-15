console.log("Script is working!");

document.getElementById('registerForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    console.log("submit");

    const messageElement = document.getElementById('auth-message');
    messageElement.innerHTML = '<div class="loading">Вход в систему...</div>';

    try {
        const response = await fetch('/auth/register', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                username: document.getElementById('username').value,
                password: document.getElementById('password').value,
                email: document.getElementById('email').value,
            })
        });

        console.log(document.getElementById('password').value);

        console.log(response.body.values);

        const data = await response.json();

        if (!response.ok) {
            const errorMsg = data.error || data.message || 'Неизвестная ошибка';
            throw new Error(errorMsg);
        }
        //window.location.href = data.redirect || '/auth/login';

    } catch (error) {
        messageElement.innerHTML = `
            <div class="error">Ошибка: ${error.message}</div>
        `;
        console.error('Ошибка входа:', error);
    }
});