const typewriter_element = document.getElementById('typewriter');

const url_input_element = $('#form input[type=url]');
const key_input_element = $('#form input[type=text]');
const submit_btn_element = $('#form button[type=submit]');

const showcase = $('#showcase');
const url_show_element = $('#showcase input[type=url]');
const copy_btn_element = $("#showcase button[type=button]");

/* ---------- Toast ---------- */

const Toast = Swal.mixin({
    toast: true,
    position: 'top',
    showConfirmButton: false,
    timer: 3000,
    showCloseButton: true
});

/* ---------- Typewriter ---------- */

const keywords = [
    "affiliators",
    "memers",
    "astronauts",
    "Elon Musk",
    "influencers",
    "YouTubers",
    "programmers",
    "developers"
];

Array.prototype.random = function () {
    return this[Math.floor(Math.random() * this.length)];
};

const typewriter = new Typewriter(typewriter_element, { loop: true });
typewriter
    .typeString(keywords.random() + '.')
    .pauseFor(1500)
    .deleteAll()
    .start();

setInterval(() => {
    typewriter
        .deleteAll()
        .typeString(keywords.random() + '.')
        .pauseFor(1500);
}, 2000);


$('#form').on('submit', function (e) {
    e.preventDefault();

    submit_btn_element.prop('disabled', true).text('Shortening...');

    const payload = {
        url: url_input_element.val().trim()
    };

    if (key_input_element.val().trim()) {
        payload.key = key_input_element.val().trim();
    }


    $.ajax({
        method: 'POST',
        url: '/shorten',
        data: JSON.stringify(payload),
        contentType: 'application/json',
    })
        .done((res) => {
            url_show_element.val(res.short_url);

            $('#form').hide();
            showcase.removeClass('d-none').show();

            Toast.fire({
                icon: 'success',
                title: 'Short URL created'
            });
        })
        .fail((xhr) => {
            let message = 'Something went wrong';

            try {
                message = JSON.parse(xhr.responseText).error;
            } catch (_) { }

            Toast.fire({
                icon: 'error',
                title: message
            });
        })
        .always(() => {
            submit_btn_element.prop('disabled', false)
                .text('Make it Tiiny');
        });
});

/* ---------- Dark / Light Mode ---------- */

const body = document.body;
const toggle = document.getElementById('mode-toggle');

function applyTheme(theme) {
    body.classList.remove('light', 'dark');
    body.classList.add(theme);
    localStorage.setItem('theme', theme);
}

const storedTheme = localStorage.getItem('theme') || 'light';
applyTheme(storedTheme);

toggle.addEventListener('click', () => {
    applyTheme(body.classList.contains('dark') ? 'light' : 'dark');
});

/* ---------- Copy ---------- */

function copyURL() {
    const input = url_show_element[0];
    input.focus();
    input.select();
    document.execCommand("copy");

    copy_btn_element.text("Copied ✅");

    setTimeout(() => {
        copy_btn_element.text("Copy");
    }, 1500);
}

/* ---------- Reset ---------- */

function resetForm() {
    showcase.hide().addClass('d-none');
    $('#form').show();

    url_input_element.val('');
    key_input_element.val('');
}
