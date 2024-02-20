package service

import (
	"html/template"

	"github.com/kkdZHC/go_blog/config"
	"github.com/kkdZHC/go_blog/dao"
	"github.com/kkdZHC/go_blog/models"
)

func GetAllIndexInfo(slug string, page, pageSize int) (*models.HomeResponse, error) {
	var categorys, err = dao.GetAllCategory() //获取category
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	var total int //总文章数
	if slug == "" {
		posts, err = dao.GetPostPage(page, pageSize) //获取post
		total = dao.CountGetAllPost()                //总文章数
	} else {
		posts, err = dao.GetPostPageBySlug(slug, page, pageSize) //获取post
		total = dao.CountGetAllPostBySlug(slug)                  //总文章数
	}

	//数据要求不是post而是postmore
	var postMores []models.PostMore
	for _, post := range posts {
		//获取postmore中缺少的两个参数
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		//post.content内容截取,否则显示过长
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[:100]
		}
		postMore := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(content),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     userName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models.DateDay(post.CreateAt),
			UpdateAt:     models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}

	pagesCount := (total-1)/10 + 1 //总页数
	var pages []int                //{1,2,3} e.g.有三页
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     postMores,
		Total:     total,
		Page:      page,
		Pages:     pages,
		PageEnd:   page != pagesCount,
	}
	return hr, nil
}
