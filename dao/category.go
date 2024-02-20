package dao

import (
	"log"

	"github.com/kkdZHC/go_blog/models"
)

func GetCategoryNameById(cId int) string {
	row := DB.QueryRow("select name from blog_category where cid=?", cId)
	if row.Err() != nil {
		log.Println("查询cId出错: ", row.Err())
		return ""
	}
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}
func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("查询blog_category出错: ", err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("取值blog_category出错: ", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}
