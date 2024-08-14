document.querySelectorAll('.notification-trigger').forEach(element => {
    element.addEventListener('click', function() {
        openModal(this.dataset.notification);
    });
});

function openModal(notificationId) {
    const notifications = {
        '1': {
            title: 'Новое сообщение',
            description: 'Вам пришло новое сообщение.',
            details: 'Это подробная информация о новом сообщении. dsfdsfsf'
        },
        '2': {
            title: 'Обновление системы',
            description: 'Система была обновлена.',
            details: 'Это подробная информация об обновлении системы.'
        },
        '3': {
            title: 'Напоминание о событии',
            description: 'Не забудьте о предстоящем событии.',
            details: 'Это подробная информация о напоминании.'
        }
    };

    const contentDiv = document.getElementById('notificationContent');
    contentDiv.innerHTML = '';

    const notification = notifications[notificationId];
    if (notification) {
        contentDiv.innerHTML = `
            <div class="notification-card bg-white shadow-md rounded-lg mb-4 p-4 cursor-pointer" onclick="showDetails('${notificationId}')">
                <div class="flex items-center">
                    <img src="https://via.placeholder.com/50" alt="Notification Icon" class="w-12 h-12 rounded-full mr-4">
                    <div>
                        <h3 class="text-lg font-semibold">${notification.title}</h3>
                        <p class="text-sm text-gray-600">${notification.description}</p>
                    </div>
                </div>
            </div>
        `;
    }

    document.getElementById('notificationModal').classList.add('modal-open');
}

function closeModal() {
    document.getElementById('notificationModal').classList.remove('modal-open');
}

function showDetails(notificationId) {
    const notifications = {
        '1': {
            title: 'Новое сообщение',
            details: 'Это подробная информация о новом сообщении.'
        },
        '2': {
            title: 'Обновление системы',
            details: 'Это подробная информация об обновлении системы.'
        },
        '3': {
            title: 'Напоминание о событии',
            details: 'Это подробная информация о напоминании.'
        }
    };

    const notification = notifications[notificationId];
    if (notification) {
        const contentDiv = document.getElementById('notificationContent');
        contentDiv.innerHTML = `
            <div class="notification-card bg-white shadow-md rounded-lg mb-4 p-4">
                <div class="flex items-center">
                    <img src="https://via.placeholder.com/50" alt="Notification Icon" class="w-12 h-12 rounded-full mr-4">
                    <div>
                        <h3 class="text-lg font-semibold">${notification.title}</h3>
                        <p class="text-sm text-gray-600">${notification.details}</p>
                    </div>
                </div>
            </div>
        `;
    }
}
