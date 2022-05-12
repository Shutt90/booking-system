function qs(node) {
    return document.querySelector(node)
}

const arrowLeft = qs('.arrow-left')
const arrowRight = qs('.arrow-right')
const calender = qs('.calender')
const date = qs('.calender-dates')
const day = qs('.calender-day')
const arrival = qs('#form-arrival')
const depart = qs('#form-depart')
const dateChildren = date.querySelectorAll('div')
let activeForm = qs('.active')

const options = [qs('#form-arrival'), qs('#form-depart')]

let calenderOpen = false;

function calenderState() {

    if(calenderOpen && this.classList.contains('active')) {
        this.classList.remove('active')
        calender.style.display = 'none';
        calenderOpen = false;
    } else if(!calenderOpen && !this.classList.contains('active')) {
        this.classList.add('active');
        calender.style.display = 'flex';
        calenderOpen = true;

    }
}

function getValue(e) {
    const selected = e.target.textContent.trim()
    if(depart.classList.contains('active')) {
        depart.value = selected 
    } else if (arrival.classList.contains('active')) {
        arrival.value = selected
    }
}

depart.addEventListener('click', calenderState, this)
arrival.addEventListener('click', calenderState, this)

dateChildren.forEach(date => {
    date.addEventListener('click', getValue)
})