'use strict';

const modal = document.querySelector('.modal');

const overlay = document.querySelector('.overlay');

const btnCloseModal = document.querySelector('.close-modal');

const btnsOpenModal = document.querySelectorAll('.show-modal');

const closeModal = () => {
  overlay.classList.add('hidden');
  modal.classList.add('hidden');
};

const openModal = () => {
  modal.classList.remove('hidden');
  overlay.classList.remove('hidden');
};

btnsOpenModal.forEach((itm) => itm.addEventListener('click', openModal));

btnCloseModal.addEventListener('click', closeModal);

modal.addEventListener('click', closeModal);

document.addEventListener('keydown', function(e) {

  if(e.key === 'Escape' && !modal.classList.contains('hidden')) {
    closeModal();
  }
});
