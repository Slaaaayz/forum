const unReadMessages = document.querySelectorAll('.unread');
const unReadMessagesCount = document.getElementById('num-of-notif');
const markAll = document.getElementById('mark-as-read');

unReadMessagesCount.innerText = unReadMessages.length;

unReadMessages.forEach((message) => {
    message.addEventListener('click', () => {
        message.classList.remove('unread');
        const newUnReadMessages = document.querySelectorAll('.unread');
        unReadMessagesCount.innerText = newUnReadMessages.length;
    });
});

markAll.addEventListener('click', () => {
    unReadMessages.forEach((message) => {
        message.classList.remove('unread');
    });
    const newUnReadMessages = document.querySelectorAll('.unread');
    unReadMessagesCount.innerHTML = newUnReadMessages.length;
});