const nav = document.querySelector('.fa-bars')
const overlay = document.querySelector('.overlay')
const content = document.querySelector('.content')
nav.addEventListener('click', () => {
  if(nav.classList.contains('active')){
    nav.classList.remove('active')
    nav.classList.add('fa-bars')
    nav.classList.remove('fa-x')
    overlay.classList.remove('active')
    content.classList.remove('active')
  } else {
    nav.classList.add('active')
    nav.classList.remove('fa-bars')
    nav.classList.add('fa-x')
    overlay.classList.add('active')
    setTimeout(function () {
      content.classList.add('active')
      overlay.style.display = 'flex'
    }, 500);
  }
})
