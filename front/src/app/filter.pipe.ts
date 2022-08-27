import { Pipe, PipeTransform } from '@angular/core';
import { Product } from './product';

@Pipe({
  name: 'filter'
})
export class FilterPipe implements PipeTransform {

  transform(products: Product[], term: string){

    if (parseInt(term) === 0){
      return products;
    }
    else {
      return products.filter((product) => {
        product.id_category === parseInt(term)
      });
    }
    

  }

}
