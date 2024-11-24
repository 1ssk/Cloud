const uploadForm = document.getElementById('uploadForm');
const fileInput = document.getElementById('fileInput');
const fileList = document.getElementById('fileList');

// Функция для загрузки файла
uploadForm.addEventListener('submit', async (event) => {
    event.preventDefault();

    const file = fileInput.files[0];
    if (!file) {
        alert('Выберите файл для загрузки');
        return;
    }

    const formData = new FormData();
    formData.append('file', file);

    try {
        const response = await fetch('/api/files', {
            method: 'POST',
            body: formData,
        });

        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(errorText || 'Ошибка загрузки файла');
        }

        alert('Файл успешно загружен');
        await loadFiles();
    } catch (error) {
        console.error('Ошибка загрузки файла:', error.message);
        alert(`Ошибка загрузки файла: ${error.message}`);
    }
});

// Функция для загрузки списка файлов
async function loadFiles() {
    try {
        const response = await fetch('/api/files/list');
        if (!response.ok) {
            throw new Error('Ошибка получения списка файлов');
        }

        const files = await response.json();
        fileList.innerHTML = '';

        files.forEach((file) => {
            const listItem = document.createElement('li');
            const link = document.createElement('a');
            link.href = `/uploads/${file}`;
            link.textContent = file;
            link.target = '_blank';
            listItem.appendChild(link);
            fileList.appendChild(listItem);
        });
    } catch (error) {
        console.error('Ошибка загрузки списка файлов:', error.message);
        alert('Ошибка загрузки списка файлов');
    }
}

// Загружаем список файлов при загрузке страницы
document.addEventListener('DOMContentLoaded', loadFiles);
