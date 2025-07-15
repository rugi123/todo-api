document.getElementById('loginForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    console.log("submit");
    // Проверяем существование элемента перед использованием
    const messageElement = document.getElementById('auth-message');

    messageElement.innerHTML = '<div class="loading">Вход в систему...</div>';

    try {
        const response = await fetch('/auth/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                username: document.getElementById('username').value,
                password: document.getElementById('password').value
            })
        });

        const data = await response.json();

        if (!response.ok) {
            // Обрабатываем разные форматы ошибок
            const errorMsg = data.error || data.message || 'Неизвестная ошибка';
            throw new Error(errorMsg);
        }

        // Сохраняем токен (если сервер его возвращает)
        if (data.token) {
            localStorage.setItem('authToken', data.token);
        }

        // Редирект после успешного входа
        window.location.href = data.redirect || '/profile';

    } catch (error) {
        messageElement.innerHTML = `
            <div class="error">Ошибка: ${error.message}</div>
        `;
        console.error('Ошибка входа:', error);
    }
});