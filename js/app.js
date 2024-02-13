const menu = document.querySelector('.menu')
const mobile_menu = document.querySelector('.toggle_menu')

let menu_toggle = false

menu.addEventListener('click', () => {
  if(menu.classList.contains('fa-bars')){
    menu.classList.remove('fa-bars')
    menu.classList.add('fa-x')
  } else {
    menu.classList.add('fa-bars')
    menu.classList.remove('fa-x')
  }
  menu.classList.add('active')
  menu_toggle ? mobile_menu.classList.remove('active') : mobile_menu.classList.add('active')
  menu_toggle = !menu_toggle
})
