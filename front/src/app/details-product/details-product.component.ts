import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Product } from '../product';
import { ProductService } from '../product.service';

@Component({
  selector: 'app-details-product',
  templateUrl: './details-product.component.html',
  styleUrls: ['./details-product.component.css']
})

export class DetailsProductComponent implements OnInit {

  id: number = 0;
  product: Product = new Product();

  constructor(private route: ActivatedRoute, private router: Router,
    private productService: ProductService) { }

  ngOnInit() {

    this.id = this.route.snapshot.params['id'];
    
    this.productService.getProduct(this.id)
      .subscribe(data => {
        this.product = data;
      });
  }

  list(){
    this.router.navigate(['products']);
  }
}
