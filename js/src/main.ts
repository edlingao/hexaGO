import '../css/style.css';
import Alpine from 'alpinejs'
import 'htmx.org';
import { Store } from './store/example';

declare global {
  interface Window {
    Alpine: typeof Alpine
    htmx: any;
  }
}

window.Alpine = Alpine
 
window.addEventListener('DOMContentLoaded', () => {
  Store();
  Alpine.start()
}, false);

