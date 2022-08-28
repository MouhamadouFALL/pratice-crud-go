import { Component, OnInit } from '@angular/core';
import { Observable, map } from "rxjs";
import { ProductService } from "../product.service";
import { Router } from '@angular/router';
import { Product } from '../product'

@Component({
  selector: 'app-list-products',
  templateUrl: './list-products.component.html',
  styleUrls: ['./list-products.component.css']
})
export class ListProductsComponent implements OnInit {

  term: any; // filter

  isContain = false

  products = new Observable<Product[]>()

  constructor(private productService: ProductService, 
    private router: Router) { }

  ngOnInit(): void {
    this.reloadData();
  }

  reloadData(){
    this.products = this.productService.getProductList()
    this.products.subscribe(isContain => this.isContain = true)

    console.log(this.products)
  }

  deleteProduct(id: number) {
    this.productService.deleteProduct(id);
    this.reloadData();
  }

  updateProduct(id: number) {
    this.router.navigate(['edit', id]);
  }

  productDetails(id: number){
    this.router.navigate(['details', id]);
  }

  search(){
    if (!(this.term == 0 || this.term == undefined)) {
      this.reloadData()
      this.products = this.products.pipe(
        map(arr => arr.filter(
          data => data.id_category == this.term
          )))
    }
    else {
      this.reloadData()
    }
  }


}
