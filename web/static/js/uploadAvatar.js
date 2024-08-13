document.addEventListener('DOMContentLoaded', function() {
    const avatarModal = document.getElementById('avatarModal');
    const openAvatarModalButton = document.getElementById('openAvatarModal');
    const avatarInput = document.getElementById('avatarInput');
    const avatarPreview = document.getElementById('avatarPreview');
    const cropButton = document.getElementById('cropButton');
    let cropper;

    // Открытие модального окна
    openAvatarModalButton.addEventListener('click', function() {
        avatarModal.classList.add('modal-open');
    });

    // Закрытие модального окна
    window.closeAvatarModal = function() {
        avatarModal.classList.remove('modal-open');
        if (cropper) {
            cropper.destroy();
        }
        avatarPreview.src = '';
        avatarInput.value = '';
    };

    // Обработка выбора изображения
    avatarInput.addEventListener('change', function(event) {
        const files = event.target.files;
        if (files && files.length > 0) {
            const reader = new FileReader();
            reader.onload = function(e) {
                avatarPreview.src = e.target.result;
                if (cropper) {
                    cropper.destroy();
                }
                cropper = new Cropper(avatarPreview, {
                    aspectRatio: 1, // Квадратный аватар
                    viewMode: 2,
                    autoCropArea: 1,
                });
            };
            reader.readAsDataURL(files[0]);
        }
    });

    // Обработка нажатия на кнопку "Сохранить"
    cropButton.addEventListener('click', function() {
        if (cropper) {
            cropper.getCroppedCanvas({
                width: 250,
                height: 250,
            }).toBlob(function(blob) {
                // Создание FormData для отправки на сервер
                const formData = new FormData();
                formData.append('avatar', blob, 'avatar.png'); // 'avatar' — это ключ, который будет использоваться на сервере

                // Выполнение PUT-запроса с использованием fetch
                fetch('/data/api/avatar', {
                    method: 'PUT',
                    body: formData
                })
                .then(response => {
                    if (response.ok) {
                        return response.json(); // Предполагаем, что сервер возвращает JSON
                    } else {
                        throw new Error('Ошибка при загрузке аватара');
                    }
                })
                .then(data => {
                    console.log('Успешно загружено:', data);
                    // Здесь можно обновить аватар на странице, если нужно
                })
                .catch(error => {
                    console.error('Ошибка:', error);
                });

                // Закрываем модальное окно после сохранения
                closeAvatarModal();
            });
        }
    });
});
